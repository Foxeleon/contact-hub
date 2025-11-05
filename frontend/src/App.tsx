import { Container, Typography, Box, CssBaseline } from '@mui/material';
import { SearchFilter } from './features/persons-list/components/SearchFilter';
import { PersonsTable } from './features/persons-list/components/PersonsTable';

function App() {
    return (
        <>
            <CssBaseline />
            <Box
                sx={{
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'center',
                    alignItems: 'center',
                    minHeight: '100vh',
                    py: 4,
                    bgcolor: 'grey.100'
                }}
            >
                <Container maxWidth="lg">
                    <Typography variant="h4" component="h1" gutterBottom>
                        Contact Hub
                    </Typography>
                    <SearchFilter />
                    <PersonsTable />
                </Container>
            </Box>
        </>
    );
}

export default App;
