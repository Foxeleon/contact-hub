import {
    Paper,
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Typography,
    TablePagination
} from '@mui/material';
import type { Person } from '../../../types/person';

// Define the props for the component
interface PersonsTableProps {
    persons: Person[];
    page: number;
    rowsPerPage: number;
    totalRows: number;
    onRowClick: (person: Person) => void;
    onPageChange: (newPage: number) => void;
    onRowsPerPageChange: (newRowsPerPage: number) => void;
}

export const PersonsTable = ({
                                 persons,
                                 page,
                                 rowsPerPage,
                                 totalRows,
                                 onRowClick,
                                 onPageChange,
                                 onRowsPerPageChange,
                             }: PersonsTableProps) => {

    // Event handlers now just call functions from props
    const handlePageChange = (_event: unknown, newPage: number) => {
        onPageChange(newPage);
    };

    const handleRowsPerPageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        onRowsPerPageChange(parseInt(event.target.value, 10));
    };

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
                            persons.map((person) => (
                                // Use a unique key if available, e.g., person.id
                                <TableRow
                                    key={`${person.firstName}-${person.lastName}`}
                                    hover
                                    sx={{ cursor: 'pointer' }}
                                    onClick={() => onRowClick(person)} // Call parent's handler on click
                                >
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
            {/* Pagination is now fully controlled by the parent component */}
            <TablePagination
                rowsPerPageOptions={[5, 10, 25]}
                component="div"
                count={totalRows}
                rowsPerPage={rowsPerPage}
                page={page}
                onPageChange={handlePageChange}
                onRowsPerPageChange={handleRowsPerPageChange}
            />
        </Paper>
    );
};