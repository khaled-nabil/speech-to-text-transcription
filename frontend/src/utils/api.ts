import { QueryClient } from '@tanstack/react-query'
import { getUserFromCookie } from 'utils/cookies.ts'

declare const __BACKEND_URL__: string

export const API_BASE_URL = __BACKEND_URL__

export const fetchAPI = async <T>(
	endpoint: string,
	options?: RequestInit
): Promise<T> => {
	const userID = getUserFromCookie()?.userId
	const response = await fetch(`${API_BASE_URL}${endpoint}`, {
		...options,
		headers: {
			...options?.headers,
			...(userID && { 'X-User-ID': userID }),
		},
	})
	if (!response.ok) {
		throw new Error('Network response was not ok')
	}
	return response.json()
}

export const queryClient = new QueryClient()
