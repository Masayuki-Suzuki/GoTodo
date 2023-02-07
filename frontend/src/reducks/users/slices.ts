import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { User } from './types'

const initialState: User = {
    id: 0,
    firstName: '',
    lastName: '',
    fullName: '',
    emailAddress: ''
}

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: {
        getUserData: (state: User) => state,
        signInAction: (state: User, action: PayloadAction<User>) => {
            const updatedData = { ...state, ...action.payload }
            return updatedData
        }
    }
})

export const { signInAction, getUserData } = userSlice.actions
export const user = userSlice.reducer
