import { getUserFromCookie } from 'utils/cookies.ts'

export const isAuthenticated = () => {
	return !!getUserFromCookie()?.userId
}
