import { Box, TextField, Button, Paper } from '@mui/material'
import { useForm } from 'react-hook-form'
import { useNavigate } from 'react-router-dom'
import { v5 as uuidv5 } from 'uuid'
import { useAppDispatch } from 'store/hooks'
import { setUser } from 'page/login/slice/userSlice'

import style from './login.module.scss'

interface LoginFormData {
	username: string
	password: string
}

const Login = () => {
	const { register, handleSubmit } = useForm<LoginFormData>()
	const navigate = useNavigate()
	const dispatch = useAppDispatch()

	const onSubmit = (data: LoginFormData) => {
		// TODO: ID should come from backend when we add authentication, for now it's a dummy'
		const userId = uuidv5(data.username, uuidv5.URL)
		dispatch(setUser({ email: data.username, userId }))
		navigate('/transcribe')
	}

	return (
		<Box className={style.container}>
			<Paper elevation={24} square={false} className={style.loginForm}>
				<form onSubmit={handleSubmit(onSubmit)}>
					<Box className={style.formContent}>
						<TextField
							label="Username"
							{...register('username', { required: true })}
							fullWidth
						/>
						<TextField
							label="Password"
							type="password"
							{...register('password', { required: true })}
							fullWidth
						/>
						<Button type="submit" variant="contained" fullWidth>
							Login
						</Button>
					</Box>
				</form>
			</Paper>
		</Box>
	)
}

export default Login
