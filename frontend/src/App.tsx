import { Container, Typography } from '@mui/material';
import { SearchFilter } from './features/persons-list/components/SearchFilter';
import { PersonsTable } from './features/persons-list/components/PersonsTable';

function App() {
    return (
        <Container sx={{ py: 4 }}>
            <Typography variant="h4" component="h1" gutterBottom>
                Contact Hub
            </Typography>
            <SearchFilter />
            <PersonsTable />
        </Container>
    );
}

export default App;