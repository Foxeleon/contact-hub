import { useState } from 'react';
import {
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    TablePagination,
    Typography
} from '@mui/material';

// Mock data representing the structure we expect from the API
const mockPersons = [
    { firstName: 'John', lastName: 'Doe', birthday: '1990-05-15' },
    { firstName: 'Jane', lastName: 'Smith', birthday: '1985-08-22' },
    { firstName: 'Peter', lastName: 'Jones', birthday: '1992-11-30' },
    { firstName: 'Mary', lastName: 'Williams', birthday: '1998-01-20' },
    { firstName: 'David', lastName: 'Brown', birthday: '1982-03-12' },
    { firstName: 'Susan', lastName: 'Miller', birthday: '2001-07-07' },
];

export const PersonsTable = () => {
    // State for pagination
    const [page, setPage] = useState(0);
    const [rowsPerPage, setRowsPerPage] = useState(5);

    // Handlers for pagination events
    const handleChangePage = (_event: unknown, newPage: number) => {
        setPage(newPage);
    };

    const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
        setRowsPerPage(parseInt(event.target.value, 10));
        setPage(0); // Reset to the first page
    };

    // Calculate the slice of data to display for the current page
    const paginatedPersons = mockPersons.slice(
        page * rowsPerPage,
        page * rowsPerPage + rowsPerPage
    );

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
                        {paginatedPersons.length > 0 ? (
                            paginatedPersons.map((person, index) => (
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
            <TablePagination
                rowsPerPageOptions={[5, 10, 25]}
                component="div"
                count={mockPersons.length} // Total number of rows
                rowsPerPage={rowsPerPage}
                page={page}
                onPageChange={handleChangePage}
                onRowsPerPageChange={handleChangeRowsPerPage}
            />
        </Paper>
    );
};