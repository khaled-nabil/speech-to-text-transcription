import { createSlice, type PayloadAction } from '@reduxjs/toolkit'
import { removeUserCookie, saveUserToCookie } from 'utils/cookies.ts'

interface UserState {
	email: string | null
	userId: string | null
}

const initialState: UserState = {
	email: null,
	userId: null,
}

const userSlice = createSlice({
	name: 'user',
	initialState,
	reducers: {
		setUser: (
			state,
			action: PayloadAction<{ email: string; userId: string }>
		) => {
			state.email = action.payload.email
			state.userId = action.payload.userId
			saveUserToCookie(action.payload)
		},
		clearUser: (state) => {
			state.email = null
			state.userId = null
			removeUserCookie()
		},
	},
})

export const { setUser, clearUser } = userSlice.actions
export default userSlice.reducer
