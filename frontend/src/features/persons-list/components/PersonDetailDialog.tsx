import { Dialog, DialogTitle, DialogContent, Typography, DialogActions, Button } from '@mui/material';
import type { Person } from '../../../types/person';

interface PersonDetailDialogProps {
    person: Person | null;
    open: boolean;
    onClose: () => void;
}

export const PersonDetailDialog = ({ person, open, onClose }: PersonDetailDialogProps) => {
    if (!person) return null;

    return (
        <Dialog open={open} onClose={onClose} fullWidth maxWidth="xs">
            <DialogTitle>{person.firstName} {person.lastName}</DialogTitle>
            <DialogContent dividers>
                <Typography gutterBottom><b>Birthday:</b> {new Date(person.birthday).toLocaleDateString()}</Typography>
                <Typography gutterBottom><b>Address:</b> {person.address || 'N/A'}</Typography>
                <Typography><b>Phone:</b> {person.phoneNumber || 'N/A'}</Typography>
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose}>Close</Button>
            </DialogActions>
        </Dialog>
    );
};