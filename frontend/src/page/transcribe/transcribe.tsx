import { Box } from '@mui/material';
import MicrophoneButton from '../../component/microphoneButton';

const Transcribe = () => {
  const handleMicrophoneClick = () => {
    // TODO: Implement microphone functionality
    console.log('Microphone clicked');
  };

  return (
    <Box
      sx={{
        height: '100vh',
        backgroundColor: '#F5F5DC', // beige color
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'flex-end',
        alignItems: 'center',
        pb: 4,
      }}
    >
      <MicrophoneButton onClick={handleMicrophoneClick} />
    </Box>
  );
};

export default Transcribe;
