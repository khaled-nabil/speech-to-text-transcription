import { BrowserRouter, Routes, Route } from 'react-router-dom'
import Home from 'page/home'
import Login from 'page/login'
import Transcribe from 'page/transcribe'
import ProtectedRoute from './ProtectedRoute'

const Router = () => {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/" element={<Home />} />
				<Route path="/login" element={<Login />} />
				<Route
					path="/transcribe"
					element={
						<ProtectedRoute>
							<Transcribe />
						</ProtectedRoute>
					}
				/>
			</Routes>
		</BrowserRouter>
	)
}

export default Router
