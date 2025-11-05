import { useState, useEffect } from 'react';
import { Container, Typography, Box, CssBaseline, CircularProgress, Alert } from '@mui/material';
import { SearchFilter } from './features/persons-list/components/SearchFilter';
import { PersonsTable } from './features/persons-list/components/PersonsTable';
import { PersonDetailDialog } from './features/persons-list/components/PersonDetailDialog';
import type { Person } from './types/person';

const API_URL = import.meta.env.VITE_API_URL_PERSONS;

// Response type from our backend
interface ApiResponse {
    Data: Person[];
    Total: number;
}

function App() {
    // Data and UI state
    const [persons, setPersons] = useState<Person[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    // Dialog state
    const [selectedPerson, setSelectedPerson] = useState<Person | null>(null);
    const [isDialogOpen, setDialogOpen] = useState(false);

    // Filtering and pagination state
    const [searchTerm, setSearchTerm] = useState('');
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);
    const [totalRows, setTotalRows] = useState(0);

    // Effect to fetch data when filters or page change
    useEffect(() => {
        const fetchPersons = async () => {
            setLoading(true);
            setError(null);

            // Build query params
            const params = new URLSearchParams();
            params.append('page', (page + 1).toString());
            params.append('pageSize', rowsPerPage.toString());
            if (searchTerm) {
                params.append('q', searchTerm);
            }
            // TODO: Add date filters to params

            try {
                const response = await fetch(`${API_URL}?${params.toString()}`);
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                const result: ApiResponse = await response.json();
                setPersons(result.Data || []);
                setTotalRows(result.Total || 0);
            } catch (e: any) {
                setError(`Failed to fetch data: ${e.message}`);
                setPersons([]);
                setTotalRows(0);
            } finally {
                setLoading(false);
            }
        };

        // Debounce fetching to avoid too many API calls while typing
        const timer = setTimeout(() => {
            fetchPersons();
        }, 500); // 500ms delay

        return () => clearTimeout(timer);

    }, [searchTerm, page, rowsPerPage]);

    // Handlers
    const handleRowClick = (person: Person) => {
        setSelectedPerson(person);
        setDialogOpen(true);
    };

    return (
        <>
            <CssBaseline />
            <Box sx={{ display: 'flex', justifyContent: 'center', py: 4, bgcolor: 'grey.100', minHeight: '100vh' }}>
                <Container maxWidth="lg">
                    <Typography variant="h4" component="h1" gutterBottom>Contact Hub</Typography>

                    <SearchFilter
                        searchTerm={searchTerm}
                        onSearchChange={setSearchTerm}
                    />

                    {loading ? (
                        <Box sx={{ display: 'flex', justifyContent: 'center', my: 4 }}><CircularProgress /></Box>
                    ) : error ? (
                        <Alert severity="error">{error}</Alert>
                    ) : (
                        <PersonsTable
                            persons={persons}
                            page={page}
                            rowsPerPage={rowsPerPage}
                            totalRows={totalRows}
                            onRowClick={handleRowClick}
                            onPageChange={setPage}
                            onRowsPerPageChange={setRowsPerPage}
                        />
                    )}

                    <PersonDetailDialog
                        person={selectedPerson}
                        open={isDialogOpen}
                        onClose={() => setDialogOpen(false)}
                    />
                </Container>
            </Box>
        </>
    );
}

export default App;