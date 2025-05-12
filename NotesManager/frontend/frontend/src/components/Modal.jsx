import React from 'react'
import PropTypes from 'prop-types'
import ReactDOM from 'react-dom'
import "./Modal.css";

export default function Modal({ isOpen, title, children, onCancel }) {
    if (!isOpen) return null
    return ReactDOM.createPortal(
        <div className="modal-overlay">
            <div className="modal-window">
                {title && <h2 className="modal-title">{title}</h2>}
                <div className="modal-content">{children}</div>
                <button className="modal-close" onClick={onCancel}>×</button>
            </div>
        </div>,
        document.body
    )
}

Modal.propTypes = {
    isOpen: PropTypes.bool.isRequired,
    title: PropTypes.string,
    children: PropTypes.node,
    onCancel: PropTypes.func.isRequired,
}
