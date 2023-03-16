import { Box, FormLabel, InputGroup } from '@chakra-ui/react'
import React from 'react'
import DatePicker from 'react-datepicker'
import { subDays } from 'date-fns'
import { Nullable } from '../../types/utils'
import 'react-datepicker/dist/react-datepicker.css'
import '../../styles/date-picker.sass'

type DatePickerPropsType = {
    action: (value: Date) => void
    value: Nullable<string>
}

const datePicker = ({ action, value }: DatePickerPropsType) => {
    return (
        <Box className="addToDo--date" mt={4}>
            <FormLabel textTransform="capitalize" color="gray.600" fontWeight={600}>
                Due Date
            </FormLabel>
            <InputGroup>
                <DatePicker
                    selected={value ? new Date(value) : new Date()}
                    onChange={action}
                    minDate={subDays(new Date(), 0)}
                    dateFormat="yyyy/MM/dd kk:mm"
                    showTimeSelect
                    isClearable
                    className="addToDo__date-picker"
                />
            </InputGroup>
        </Box>
    )
}

export default datePicker
