import { Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography } from '@mui/material';

export const PersonsTable = () => {
    return (
        <TableContainer component={Paper}>
            <Table>
                <TableHead>
                    <TableRow>
                        <TableCell>Name</TableCell>
                        <TableCell>Surname</TableCell>
                        <TableCell>Date of birth</TableCell>
                    </TableRow>
                </TableHead>
                <TableBody>
                    {/* TODO: Replace with real data */}
                    <TableRow>
                        <TableCell colSpan={3} align="center">
                            <Typography>There is no data to display</Typography>
                        </TableCell>
                    </TableRow>
                </TableBody>
            </Table>
            {/* TODO: Add Pagination */}
        </TableContainer>
    );
};