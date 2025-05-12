import React, { useState, useRef, useEffect } from "react";
import PropTypes from "prop-types";
import "./ListCard.css";

export default function ListCard({ list, onEdit, onDelete }) {
  const [menuOpen, setMenuOpen] = useState(false);
  const ref = useRef(null);

  useEffect(() => {
    const handleClick = (e) => {
      if (ref.current && !ref.current.contains(e.target)) {
        setMenuOpen(false);
      }
    };
    document.addEventListener("mousedown", handleClick);
    return () => document.removeEventListener("mousedown", handleClick);
  }, []);

  const handleEdit = () => {
    setMenuOpen(false);
    onEdit();
    console.log('card')
  };

  const handleDelete = () => {
    setMenuOpen(false);
    onDelete();
  };

  return (
    <div className="card" ref={ref}>
      <div className="card__header">
        <h3 className="card__title">{list.title}</h3>
        <button
          className="card__menu-btn"
          onClick={() => setMenuOpen((o) => !o)}
          aria-label="Open actions menu"
        >
          ⋮
        </button>
        {menuOpen && (
          <div className="card__menu" role="menu">
            <button className="card__menu-item" onClick={handleEdit}>
              Edit
            </button>
            <button className="card__menu-item" onClick={handleDelete}>
              Delete
            </button>
          </div>
        )}
      </div>

      {list.description && (
        <div className="card__body">
          <p className="card__desc">{list.description}</p>
        </div>
      )}

      <div className="card__meta">
        <span className="card__date">
          Created: {new Date(list.created_at).toLocaleString()}
        </span>
      </div>

      <div className="card__footer">
        <button
          className="btn btn--primary card__open-btn"
          onClick={() => (window.location.href = `/lists/${list.id}`)}
        >
          Open
        </button>
      </div>
    </div>
  );
}

ListCard.propTypes = {
  list: PropTypes.shape({
    id: PropTypes.number.isRequired,
    title: PropTypes.string.isRequired,
    description: PropTypes.string,
    created_at: PropTypes.string.isRequired,
  }).isRequired,
  onEdit: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
};
