import { type ReactNode } from 'react'
import { Navigate } from 'react-router-dom'
import { useAppSelector } from 'store/hooks.ts'

const ProtectedRoute = ({ children }: { children: ReactNode }) => {
	const userID = useAppSelector(({ user }) => user.userId)

	if (!userID) {
		return <Navigate to="/login" replace />
	}

	return children
}

export default ProtectedRoute
