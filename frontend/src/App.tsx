import { useState, useEffect } from 'react';
import { Container, Typography, Box, CssBaseline, CircularProgress, Alert } from '@mui/material';
import { SearchFilter } from './features/persons-list/components/SearchFilter';
import { PersonsTable } from './features/persons-list/components/PersonsTable';
import { PersonDetailDialog } from './features/persons-list/components/PersonDetailDialog';
import type { Person } from './types/person';

const API_URL = import.meta.env.VITE_API_URL_PERSONS;

// The response type from our Go backend
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
    const [birthdayFrom, setBirthdayFrom] = useState<Date | null>(null);
    const [birthdayTo, setBirthdayTo] = useState<Date | null>(null);
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(10);
    const [totalRows, setTotalRows] = useState(0);

    // Effect to fetch data when any filter or page changes
    useEffect(() => {
        const fetchPersons = async () => {
            if (!API_URL) {
                setError('API URL is not configured. Please check your .env file.');
                setLoading(false);
                return;
            }

            setLoading(true);
            setError(null);

            // Build query params from all filter states
            const params = new URLSearchParams();
            params.append('page', (page + 1).toString());
            params.append('pageSize', rowsPerPage.toString());
            if (searchTerm) {
                params.append('q', searchTerm);
            }
            // Add date filters to params if they are selected, in ISO format
            if (birthdayFrom) {
                params.append('birthdayFrom', birthdayFrom.toISOString());
            }
            if (birthdayTo) {
                // Set to the end of the day to ensure the selected day is included in the range
                const endOfDay = new Date(birthdayTo);
                endOfDay.setHours(23, 59, 59, 999);
                params.append('birthdayTo', endOfDay.toISOString());
            }

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

        // Debounce fetching to avoid excessive API calls while typing
        const debounce = setTimeout(() => {
            fetchPersons();
        }, 500);

        return () => clearTimeout(debounce);

        // Add new date states to the dependency array to re-trigger the effect
    }, [searchTerm, birthdayFrom, birthdayTo, page, rowsPerPage]);

    // Handlers
    const handleRowClick = (person: Person) => {
        setSelectedPerson(person);
        setDialogOpen(true);
    };

    // Handler to reset page to 0 when rowsPerPage changes
    const handleRowsPerPageChange = (newRowsPerPage: number) => {
        setRowsPerPage(newRowsPerPage);
        setPage(0);
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
                        birthdayFrom={birthdayFrom}
                        onBirthdayFromChange={setBirthdayFrom}
                        birthdayTo={birthdayTo}
                        onBirthdayToChange={setBirthdayTo}
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
                            onRowsPerPageChange={handleRowsPerPageChange}
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