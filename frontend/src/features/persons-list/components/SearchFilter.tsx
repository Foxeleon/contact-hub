import { Box, TextField, Paper } from '@mui/material';

export const SearchFilter = () => {
    return (
        <Paper sx={{ p: 2, mb: 3 }}>
            <Box component="form" noValidate autoComplete="off">
                <TextField label="Поиск по тексту..." variant="outlined" fullWidth />
                {/* TODO: Add filters by date of birth */}
            </Box>
        </Paper>
    );
};