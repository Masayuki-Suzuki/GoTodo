import React, { useState } from 'react'
import axios from 'axios'
import { Box, Button, Spinner } from '@chakra-ui/react'
import { useFormik } from 'formik'
import { CardTitle } from '../atoms/CardTitle'
import CommonInputField from '../molecules/CommonInputField'
import ConfirmField from '../molecules/ConfirmField'
import { signUpSchema } from '../../libs/validationSchema'

type InitialValues = {
    emailAddress: string
    password: string
    confirmPass: string
    firstName: string
    lastName: string
}

const SignUpCard = () => {
    const [loading, setLoading] = useState<boolean>(false)

    const { errors, handleChange, handleSubmit, touched, values } = useFormik<InitialValues>({
        initialValues: {
            emailAddress: '',
            password: '',
            confirmPass: '',
            firstName: '',
            lastName: ''
        },
        validationSchema: signUpSchema,
        async onSubmit(values) {
            try {
                setLoading(true)
                const res = await axios.post('http://localhost:4000/api/admin/register', values)
                console.log(res)
                setLoading(false)
            } catch (e) {
                console.error(e)
            }
        }
    })

    return (
        <Box w={400} border="1px" borderColor="gray.100" boxShadow="sm" p="6" borderRadius={6}>
            <CardTitle title="Sign Up" />
            <form onSubmit={handleSubmit}>
                <CommonInputField
                    onInputAction={handleChange}
                    name="firstName"
                    inputType="text"
                    label="First Name"
                    placeholder="John"
                    value={values.firstName}
                    error={errors.firstName && touched.firstName ? errors.firstName : null}
                    mt={6}
                    isRequired
                />
                <CommonInputField
                    inputType="text"
                    label="Last Name"
                    placeholder="Doe"
                    name="lastName"
                    onInputAction={handleChange}
                    value={values.lastName}
                    error={errors.lastName && touched.lastName ? errors.lastName : null}
                    mt={4}
                    isRequired
                />
                <CommonInputField
                    onInputAction={handleChange}
                    inputType="email"
                    label="email"
                    placeholder="example@example.com"
                    name="emailAddress"
                    value={values.emailAddress}
                    error={errors.emailAddress && touched.emailAddress ? errors.emailAddress : null}
                    mt={4}
                    isRequired
                />
                <ConfirmField label="confirm email address" mt={4} email={values.emailAddress} />
                <CommonInputField
                    inputType="password"
                    label="password"
                    placeholder=""
                    name="password"
                    onInputAction={handleChange}
                    value={values.password}
                    error={errors.password && touched.password ? errors.password : null}
                    mt={4}
                    isRequired
                />
                <CommonInputField
                    inputType="password"
                    label="confirm password"
                    placeholder=""
                    name="confirmPass"
                    onInputAction={handleChange}
                    value={values.confirmPass}
                    error={errors.confirmPass && touched.confirmPass ? errors.confirmPass : null}
                    mt={4}
                    isRequired
                />
                <Box mt={6} textAlign="center">
                    {loading ? (
                        <Spinner thickness="4px" speed="0.65s" emptyColor="gray.200" color="blue.500" size="xl" />
                    ) : (
                        <Button type="submit" w={120} bg="#19448e" color="white" fontWeight={700} boxShadow="base">
                            Sign Up
                        </Button>
                    )}
                </Box>
            </form>
        </Box>
    )
}

export default SignUpCard
