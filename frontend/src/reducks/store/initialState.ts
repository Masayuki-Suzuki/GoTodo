import { RootState } from '../../types/store'

const initialState: RootState = {
    user: {
        id: 0,
        firstName: '',
        lastName: '',
        fullName: '',
        emailAddress: ''
    },
    fetchStatus: {
        isLoading: false,
        error: null
    },
    modalState: {
        isOpen: false
    }
}

export default initialState
