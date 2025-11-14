import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Layout from 'component/layout/layout'
import Home from 'page/home'
import Login from 'page/login'
import Transcribe from 'page/transcribe'
import ProtectedRoute from './ProtectedRoute'

const Router = () => (
	<BrowserRouter>
		<Routes>
			<Route element={<Layout />}>
				<Route path="/" element={<Home />} />
				<Route path="/login" element={<Login />} />
				<Route path="/transcribe" element={<ProtectedRoute><Transcribe /></ProtectedRoute>} />
				<Route path="/history" element={<ProtectedRoute><></></ProtectedRoute>} />
			</Route>
		</Routes>
	</BrowserRouter>
)

export default Router
