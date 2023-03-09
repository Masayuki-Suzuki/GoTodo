import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import { useFormik } from 'formik'
import { Box, Button, ModalBody, ModalFooter, Spinner, Text } from '@chakra-ui/react'
import { getFetchError, getFetchStatus } from '../../reducks/fetchStatus/selectors'
import { addToDoSchema } from '../../libs/validationSchema'
import { resetFetchError } from '../../reducks/fetchStatus/slices'
import CommonInputField from '../molecules/CommonInputField'
import ToDoEditor from '../molecules/ToDoEditor'

type InitialValues = {
    title: string
    description: string
    dueDate: string
}

const AddToDoForm = () => {
    const isLoading = useSelector(getFetchStatus)
    const isFetchError = useSelector(getFetchError)
    const dispatch = useDispatch()

    const { errors, handleChange, handleSubmit, touched, values, setFieldValue } = useFormik<InitialValues>({
        initialValues: {
            title: '',
            description: '',
            dueDate: ''
        },
        validationSchema: addToDoSchema,
        async onSubmit(values) {
            dispatch(resetFetchError())
            console.log(values)
            // await login(values)(dispatch)
        }
    })

    const onEditorChange = (value: string) => setFieldValue('description', value)

    return (
        <>
            <form onSubmit={handleSubmit}>
                <ModalBody>
                    <CommonInputField
                        onInputAction={handleChange}
                        inputType="text"
                        label="Title"
                        placeholder="ToDo Title"
                        isRequired
                        name="title"
                        value={values.title}
                        error={errors.title && touched.title ? errors.title : null}
                        mt={0}
                    />
                    <ToDoEditor onInputAction={onEditorChange} value={values.description} />
                    {isFetchError ? (
                        <Text fontSize={16} textAlign="center" textColor="red.500" fontWeight={700} mt={4}>
                            {isFetchError}
                        </Text>
                    ) : null}
                </ModalBody>
                <ModalFooter>
                    <Box mt={6} textAlign="center">
                        {isLoading ? (
                            <Spinner thickness="4px" speed="0.65s" emptyColor="gray.200" color="blue.500" size="xl" />
                        ) : (
                            <Button type="submit" w={120} bg="#19448e" color="white" fontWeight={700} boxShadow="base">
                                Add
                            </Button>
                        )}
                    </Box>
                </ModalFooter>
            </form>
        </>
    )
}

export default AddToDoForm
