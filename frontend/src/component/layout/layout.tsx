import { Container, Box, BottomNavigation, BottomNavigationAction } from '@mui/material'
import RecordIcon from '@mui/icons-material/SettingsVoice'
import HistoryIcon from '@mui/icons-material/HistoryEdu'
import LogoutIcon from '@mui/icons-material/Logout'
import { Outlet, useNavigate } from 'react-router-dom'
import style from './layout.module.scss'
import { isAuthenticated } from 'utils/auth.ts'
import { useAppDispatch } from 'store/hooks.ts'
import { clearUser } from 'page/login/slice/userSlice.ts'

const Layout = () => {
	const dispatch = useAppDispatch()
	const navigate = useNavigate()

	return (
		<div className={style.page}>
			<Container maxWidth="sm" className={style.wrapper}>
				<Box component="main" className={style.main}>
					<Outlet />
					{isAuthenticated() && (
						<BottomNavigation showLabels>
							<BottomNavigationAction icon={<RecordIcon />} onClick={() => navigate('/transcribe')} />
							<BottomNavigationAction icon={<HistoryIcon />} onClick={() => navigate('/history')} />
							<BottomNavigationAction icon={<LogoutIcon />} onClick={() => dispatch(clearUser())} />
						</BottomNavigation>
					)}
				</Box>
			</Container>
		</div>
	)
}

export default Layout
