import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { getFetchError, getFetchStatus } from '~/reducks/fetchStatus/selectors'
import { useFormik } from 'formik'
import { loginSchema } from '~/libs/validationSchema'
import { resetFetchError } from '~/reducks/fetchStatus/slices'
import { login } from '~/reducks/users/operations'
import CommonInputField from '~/components/molecules/CommonInputField'
import { Box, Button, Spinner, Text } from '@chakra-ui/react'

const AddToDoForm = () => {
    const isLoading = useSelector(getFetchStatus)
    const isFetchError = useSelector(getFetchError)
    const dispatch = useDispatch()

    const { errors, handleChange, handleSubmit, touched, values } = useFormik<InitialValues>({
        initialValues: {
            emailAddress: '',
            password: ''
        },
        validationSchema: loginSchema,
        async onSubmit(values) {
            dispatch(resetFetchError())
            // await login(values)(dispatch)
        }
    })

    return (
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
            {isFetchError ? (
                <Text fontSize={16} textAlign="center" textColor="red.500" fontWeight={700} mt={4}>
                    {isFetchError}
                </Text>
            ) : null}
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
    )
}

export default AddToDoForm
