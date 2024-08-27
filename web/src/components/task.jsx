import React, { useState } from "react";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faTrashAlt } from '@fortawesome/free-regular-svg-icons'

export default function TaskComponent({ task, onUpdate, onComplete }) {
    const [value, setValue] = useState(task.text)

    function handleChange(e) {
        setValue(e.target.value)
    }

    function handleUpdate() {
        onUpdate({"id": task.id, "text": value});
    }

    function handleClick(e) {
        onComplete(task.id)
    }

    return (
        <div className="w-full flex flex-row gap-2 drop-shadow-sm border-b-2">
            <input className="w-full text-lg outline-none" value={value} onChange={handleChange} onBlur={handleUpdate} onKeyDown={(e) => {if (e.key === 'Enter') handleUpdate();}}/>
            <button className="p-2 px-4 rounded text-red-600 hover:bg-slate-100" onClick={handleClick}>
                <FontAwesomeIcon icon={faTrashAlt}></FontAwesomeIcon>
            </button>
        </div>
    )
}