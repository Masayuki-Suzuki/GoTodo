import * as Yup from 'yup'
import axios from 'axios'

type EmailExistanceAPIResponse = {
    error: unknown
    is_used: boolean
}

export const signUpSchema = Yup.object({
    firstName: Yup.string().required('First name is required.'),
    lastName: Yup.string().required('Last name is required'),
    emailAddress: Yup.string()
        .email('Not a valid email address.')
        .required('Email address is required.')
        .test('unique', 'This email address has already taken.', async value => {
            const { data } = await axios.post<EmailExistanceAPIResponse>(
                'http://localhost:4000/api/admin/email-existence',
                {
                    emailAddress: value
                }
            )
            return !data.is_used
        }),
    password: Yup.string().min(6).required('Password is required.'),
    confirmPass: Yup.string()
        .min(6)
        .oneOf([Yup.ref('password'), null], `Passwords don't match.`)
        .required('Confirm Password is required.')
})

export const loginSchema = Yup.object({
    emailAddress: Yup.string().email('Not a valid email address.').required('Email address is required.'),
    password: Yup.string().required('Password is required.')
})
