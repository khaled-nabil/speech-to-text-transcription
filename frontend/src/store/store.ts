import { configureStore } from '@reduxjs/toolkit'
import userReducer from 'page/login/slice/userSlice'
import transcriptionReducer from 'page/transcribe/slice/transcriptionSlice'
import { getUserFromCookie } from 'utils/cookies'

const preloadedState = {
	user: {
		email: getUserFromCookie()?.email || null,
		userId: getUserFromCookie()?.userId || null,
	},
}

export const store = configureStore({
	reducer: {
		user: userReducer,
		transcription: transcriptionReducer,
	},
	preloadedState,
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
