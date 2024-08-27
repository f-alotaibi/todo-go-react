import React, { useState } from "react";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faPlusSquare } from '@fortawesome/free-regular-svg-icons'

export default function TaskInputComponent({ onUpdate }) {
    const [value, setValue] = useState("")
    function handleChange(e) {
        setValue(e.target.value)
    }
    function handleClick(e) {
        if (value == "") {
            return
        }
        onUpdate(value)
        setValue("")
    }
    return (
        <div className="w-full flex flex-row gap-2 drop-shadow-sm border-b-2">
            <input placeholder='Write your ideas here' className="w-full text-lg outline-none" value={value} onChange={handleChange}/>
            <button className="p-2 px-4 rounded text-green-600 hover:bg-slate-100" onClick={handleClick}>
                <FontAwesomeIcon icon={faPlusSquare}></FontAwesomeIcon>
            </button>
        </div>
    )
}