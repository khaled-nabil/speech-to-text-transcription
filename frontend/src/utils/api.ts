import { QueryClient } from '@tanstack/react-query'

declare const __BACKEND_URL__: string

const API_BASE_URL = __BACKEND_URL__

export const fetchAPI = async <T>(
	endpoint: string,
	options?: RequestInit
): Promise<T> => {
	const response = await fetch(`${API_BASE_URL}${endpoint}`, options)
	if (!response.ok) {
		throw new Error('Network response was not ok')
	}
	return response.json()
}


export const queryClient = new QueryClient()
