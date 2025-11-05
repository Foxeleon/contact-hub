import { Box, TextField, Paper } from '@mui/material';
import { DatePicker, LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import type { ChangeEvent } from "react";

// Define the complete props interface for the component
interface SearchFilterProps {
    searchTerm: string;
    onSearchChange: (newSearchTerm: string) => void;
    // Add props for date filters
    birthdayFrom: Date | null;
    onBirthdayFromChange: (date: Date | null) => void;
    birthdayTo: Date | null;
    onBirthdayToChange: (date: Date | null) => void;
}

// The component remains a "dumb" presentational component
export const SearchFilter = ({
                                 searchTerm,
                                 onSearchChange,
                                 birthdayFrom,
                                 onBirthdayFromChange,
                                 birthdayTo,
                                 onBirthdayToChange
                             }: SearchFilterProps) => {

    const handleSearchChange = (event: ChangeEvent<HTMLInputElement>) => {
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
                        flexDirection: { xs: 'column', sm: 'row' }
                    }}
                >
                    <TextField
                        label="Search by name..."
                        variant="outlined"
                        fullWidth
                        value={searchTerm}
                        onChange={handleSearchChange}
                    />
                    <Box sx={{ width: { xs: '100%', sm: 180 } }}>
                        <DatePicker
                            label="Birthday from"
                            value={birthdayFrom}
                            onChange={onBirthdayFromChange}
                            slotProps={{
                                textField: { fullWidth: true }
                            }}
                        />
                    </Box>
                    <Box sx={{ width: { xs: '100%', sm: 180 } }}>
                        <DatePicker
                            label="Birthday to"
                            value={birthdayTo}
                            onChange={onBirthdayToChange}
                            slotProps={{
                                textField: { fullWidth: true }
                            }}
                        />
                    </Box>
                </Box>
            </LocalizationProvider>
        </Paper>
    );
};