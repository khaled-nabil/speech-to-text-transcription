import { createSlice, type PayloadAction } from '@reduxjs/toolkit';

interface UserState {
  email: string | null;
  userId: string | null;
}

const initialState: UserState = {
  email: null,
  userId: null,
};

const userSlice = createSlice({
  name: 'user',
  initialState,
  reducers: {
    setUser: (state, action: PayloadAction<{ email: string; userId: string }>) => {
      state.email = action.payload.email;
      state.userId = action.payload.userId;
    },
    clearUser: (state) => {
      state.email = null;
      state.userId = null;
    },
  },
});

export const { setUser, clearUser } = userSlice.actions;
export default userSlice.reducer;
