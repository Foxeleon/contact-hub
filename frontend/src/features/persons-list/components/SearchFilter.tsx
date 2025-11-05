import { Box, TextField, Paper } from '@mui/material';
import { DatePicker, LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import type { ChangeEvent } from "react";

// Define the props that the component will receive
interface SearchFilterProps {
    searchTerm: string;
    onSearchChange: (newSearchTerm: string) => void;
    // TODO: Add props for date filters (e.g., birthdayFrom, onBirthdayFromChange)
}

// The component is now a "dumb" presentational component
export const SearchFilter = ({ searchTerm, onSearchChange }: SearchFilterProps) => {

    const handleSearchChange = (event: ChangeEvent<HTMLInputElement>) => {
        // It doesn't set its own state, it just calls the function passed from the parent
        onSearchChange(event.target.value);
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
                        fullWidth
                        value={searchTerm} // Value comes from props
                        onChange={handleSearchChange} // Handler calls a function from props
                    />
                    <Box sx={{ width: { xs: '100%', sm: 180 } }}>
                        <DatePicker
                            label="Birthday from"
                            slotProps={{
                                textField: { fullWidth: true }
                            }}
                            // TODO: Connect to state and handlers
                        />
                    </Box>
                    <Box sx={{ width: { xs: '100%', sm: 180 } }}>
                        <DatePicker
                            label="Birthday to"
                            slotProps={{
                                textField: { fullWidth: true }
                            }}
                            // TODO: Connect to state and handlers
                        />
                    </Box>
                </Box>
            </LocalizationProvider>
        </Paper>
    );
};