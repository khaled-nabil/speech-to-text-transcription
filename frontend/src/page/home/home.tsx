import {Box, Button} from '@mui/material';
import {useNavigate} from 'react-router-dom';

import style from './home.module.scss';

const Home = () => {
    const navigate = useNavigate();

    return (
        <Box className={style.container}>
			<div className={style.fader}/>
            <Button
                variant="contained"
                size="large"
				color="primary"
                onClick={() => navigate('/login')}
                className={style.cta}
            >
                Start
            </Button>
        </Box>
    );
};

export default Home;
