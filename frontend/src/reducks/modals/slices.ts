import { createSlice, PayloadAction } from '@reduxjs/toolkit'
import { ModalState } from './types'

const initialState: ModalState = {
    isOpen: false
}

export const modalSlice = createSlice({
    name: 'modal',
    initialState,
    reducers: {
        setOpenStatus: (state: ModalState, action: PayloadAction<ModalState>) => {
            const updatedData = { ...state, ...action.payload }
            return updatedData
        }
    }
})

export const { setOpenStatus } = modalSlice.actions
export const modalState = modalSlice.reducer
