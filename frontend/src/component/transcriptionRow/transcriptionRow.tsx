import { ListItem, ListItemText, Box, Typography } from '@mui/material';
import {type Transcription, TranscriptionStatus } from '../../page/transcription/transcription.tsx';

interface TranscriptionRowProps {
  transcription: Transcription;
}

const getStatusColor = (status: TranscriptionStatus) => {
  switch (status) {
    case TranscriptionStatus.PENDING:
      return '#FFF59D'; // yellow
    case TranscriptionStatus.ERROR:
      return '#EF5350'; // red
    case TranscriptionStatus.SUCCESS:
      return '#81C784'; // green
    default:
      return 'transparent';
  }
};

const TranscriptionRow = ({ transcription }: TranscriptionRowProps) => {
  const formattedDate = new Date(transcription.uploadDate).toLocaleString();

  return (
    <ListItem
      sx={{
        backgroundColor: getStatusColor(transcription.status),
        mb: 1,
        borderRadius: 1,
      }}
    >
      <Box sx={{ width: '100%' }}>
        <Typography variant="caption" color="text.secondary">
          {formattedDate}
        </Typography>
        <ListItemText
          primary={transcription.transcriptText}
          secondary={`Status: ${transcription.status}`}
        />
      </Box>
    </ListItem>
  );
};

export default TranscriptionRow;
