import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { Link as ReachLink } from 'react-router-dom'
import { useFormik } from 'formik'
import { Box, Button, Link, Spinner, Text } from '@chakra-ui/react'
import { CardTitle } from '../atoms/CardTitle'
import CommonInputField from '../molecules/CommonInputField'
import { signUpSchema } from '../../libs/validationSchema'
import { getFetchStatus } from '../../reducks/fetchStatus/selectors'
import { login } from '../../reducks/users/operations'

type InitialValues = {
    emailAddress: string
    password: string
}

const LogInCard = () => {
    const isLoading = useSelector(getFetchStatus)
    const dispatch = useDispatch()

    const { errors, handleChange, handleSubmit, touched, values } = useFormik<InitialValues>({
        initialValues: {
            emailAddress: '',
            password: ''
        },
        validationSchema: signUpSchema,
        async onSubmit(values) {
            try {
                await login(values)(dispatch)
            } catch (e) {
                console.error(e)
            }
        }
    })

    // ToDo: Need to fix email validation since it's for Sign Up!! (existence check runs)

    return (
        <Box w={400} border="1px" borderColor="gray.100" boxShadow="sm" p="6" borderRadius={6}>
            <CardTitle title="Sign Up" />
            <form onSubmit={handleSubmit}>
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
                <Box mt={6} textAlign="center">
                    {isLoading ? (
                        <Spinner thickness="4px" speed="0.65s" emptyColor="gray.200" color="blue.500" size="xl" />
                    ) : (
                        <Button type="submit" w={120} bg="#19448e" color="white" fontWeight={700} boxShadow="base">
                            Log In
                        </Button>
                    )}
                </Box>
            </form>
            <Text fontSize={14} textAlign="center" textColor="gray.600" mt={4}>
                {`Haven't you got an account? `}
                <Link as={ReachLink} to="/sign-up" textDecoration="underline">
                    Sign up here.
                </Link>
            </Text>
        </Box>
    )
}

export default LogInCard
