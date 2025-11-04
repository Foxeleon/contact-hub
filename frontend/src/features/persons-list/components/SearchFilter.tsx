import { useState } from 'react';
import { Box, TextField, Paper } from '@mui/material';
import { DatePicker, LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';

export const SearchFilter = () => {
    const [searchTerm, setSearchTerm] = useState('');

    // Handler for the text search input change
    const handleSearchChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSearchTerm(event.target.value);
        // For now, just log to console to confirm it's working
        console.log('Search Term:', event.target.value);
        // TODO: This will later trigger an API call based on the search term
    };

    return (
        <Paper sx={{ p: 2, mb: 3 }}>
            <LocalizationProvider dateAdapter={AdapterDateFns}>
                <Box
                    sx={{
                        display: 'flex',
                        gap: 2, // Defines the space between flex items
                        alignItems: 'center',
                        // Responsive layout: column on extra-small screens, row on small and up
                        flexDirection: { xs: 'column', sm: 'row' }
                    }}
                >
                    <TextField
                        label="Search by name..."
                        variant="outlined"
                        fullWidth // Allows the text field to take up available space
                        value={searchTerm}
                        onChange={handleSearchChange}
                    />
                    <DatePicker
                        label="Birthday from"
                        sx={{ minWidth: { sm: 180 } }} // Ensure consistent width on larger screens
                        // TODO: Implement state and onChange handler for this date picker
                    />
                    <DatePicker
                        label="Birthday to"
                        sx={{ minWidth: { sm: 180 } }} // Ensure consistent width on larger screens
                        // TODO: Implement state and onChange handler for this date picker
                    />
                </Box>
            </LocalizationProvider>
        </Paper>
    );
};