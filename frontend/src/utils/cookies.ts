import Cookies from 'js-cookie'

const AUTH_COOKIE_NAME = 'user_auth'

export interface UserAuth {
	email: string
	userId: string
}

export const saveUserToCookie = (user: UserAuth) => {
	Cookies.set(AUTH_COOKIE_NAME, JSON.stringify(user), { expires: 7 })
}

export const getUserFromCookie = (): UserAuth | null => {
	const cookie = Cookies.get(AUTH_COOKIE_NAME)
	if (!cookie) return null

	try {
		return JSON.parse(cookie)
	} catch {
		return null
	}
}

export const removeUserCookie = () => {
	Cookies.remove(AUTH_COOKIE_NAME)
}
