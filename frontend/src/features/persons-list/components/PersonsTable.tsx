import { useState, useEffect } from 'react';
import {
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Typography,
    CircularProgress,
    Alert,
    Box
} from '@mui/material';
import type { Person } from '../../../types/person'; // Import the type
// Import the type

// The API endpoint address
const API_URL = import.meta.env.API_URL;

export const PersonsTable = () => {
    // State for managing data, loading status, and errors
    const [persons, setPersons] = useState<Person[]>([]);
    const [loading, setLoading] = useState<boolean>(true);
    const [error, setError] = useState<string | null>(null);

    // Fetch data from the backend when the component mounts
    useEffect(() => {
        const fetchPersons = async () => {
            try {
                setLoading(true);
                setError(null);
                const response = await fetch(API_URL);
                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }
                const data: Person[] = await response.json();
                setPersons(data);
            } catch (e: any) {
                setError(`Failed to fetch data: ${e.message}`);
            } finally {
                setLoading(false);
            }
        };

        fetchPersons().then(() => true);
    }, []); // The empty dependency array ensures this runs only once on mount

    if (loading) {
        return (
            <Box sx={{ display: 'flex', justifyContent: 'center', my: 4 }}>
                <CircularProgress />
            </Box>
        );
    }

    if (error) {
        return <Alert severity="error">{error}</Alert>;
    }

    return (
        <Paper>
            <TableContainer>
                <Table>
                    <TableHead>
                        <TableRow>
                            <TableCell>Name</TableCell>
                            <TableCell>Surname</TableCell>
                            <TableCell>Date of birth</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {persons.length > 0 ? (
                            persons.map((person, index) => (
                                <TableRow key={index} hover sx={{ cursor: 'pointer' }}>
                                    <TableCell>{person.firstName}</TableCell>
                                    <TableCell>{person.lastName}</TableCell>
                                    <TableCell>{new Date(person.birthday).toLocaleDateString()}</TableCell>
                                </TableRow>
                            ))
                        ) : (
                            <TableRow>
                                <TableCell colSpan={3} align="center">
                                    <Typography>There is no data to display</Typography>
                                </TableCell>
                            </TableRow>
                        )}
                    </TableBody>
                </Table>
            </TableContainer>
            {/* TODO: Server-side pagination will be implemented here */}
        </Paper>
    );
};