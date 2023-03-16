import React from 'react'
import { Box, FormLabel } from '@chakra-ui/react'
import ReactQuill from 'react-quill'
import 'react-quill/dist/quill.snow.css'
import '../../styles/quill.sass'

type ToDoEditorProps = {
    value: string
    onInputAction: (value: string) => void
}

const ToDoEditor = ({ value, onInputAction }: ToDoEditorProps) => {
    return (
        <Box className="addToDo--description" mt={4}>
            <FormLabel textTransform="capitalize" color="gray.600" fontWeight={600}>
                Description
            </FormLabel>
            <ReactQuill theme="snow" value={value} onChange={onInputAction} />
        </Box>
    )
}

export default ToDoEditor
