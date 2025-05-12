import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import api from "../shared/api";
import Modal from "../components/Modal";
import "./NotePage.css";

export default function NotePage() {
  const { listId, itemId } = useParams();
  const navigate = useNavigate();

  const [note, setNote] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [pinned, setPinned] = useState(false);

  const [confirmDelete, setConfirmDelete] = useState(false);

  useEffect(() => {
    (async () => {
      setLoading(true);
      try {
        const { data } = await api.get(`/api/items/${itemId}`);
        const n = data.data ?? data;
        setNote(n);
        setTitle(n.title);
        setContent(n.content);
        setPinned(n.pinned);
      } catch (e) {
        setError(e.message);
      } finally {
        setLoading(false);
      }
    })();
  }, [itemId]);

  const handleSave = async () => {
    try {
      await api.put(`/api/items/${itemId}`, { title, content, pinned });
      navigate(`/lists/${listId}`);
    } catch (e) {
      setError(e.message);
    }
  };

  const handleDelete = async () => {
    try {
      await api.delete(`/api/items/${itemId}`);
      navigate(`/lists/${listId}`);
    } catch (e) {
      setError(e.message);
    }
  };

  if (loading) return <p>Loading…</p>;
  if (error) return <p className="text text--error">{error}</p>;
  if (!note) return <p className="text text--muted">Note not found.</p>;

  return (
    <div className="note-page">
      <header className="edit-header">
        <button
          className="btn btn--link"
          onClick={() => navigate(`/lists/${listId}`)}
        >
          ← Back
        </button>
        <h2 className="edit-title">Edit note</h2>
      </header>

      <main className="edit-main">
        <div className="edit-card">
          <h3 className="edit-card__caption">Note details</h3>

          <div className="note-edit-form">
            <input
              className="input-title"
              placeholder="Title"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
            />

            <textarea
              className="textarea-content"
              placeholder="Content"
              rows={8}
              value={content}
              onChange={(e) => setContent(e.target.value)}
            />

            <label className="form-switch">
              <input
                type="checkbox"
                checked={pinned}
                onChange={(e) => setPinned(e.target.checked)}
              />
              <span className="slider" />
              <span className="switch-label">Pinned</span>
            </label>

            <div className="form-actions">
              <button className="btn btn--primary" onClick={handleSave}>
                Save changes
              </button>
              <button
                className="btn btn--error"
                onClick={() => setConfirmDelete(true)}
              >
                Delete
              </button>
            </div>
          </div>
        </div>
      </main>

      <Modal
        isOpen={confirmDelete}
        title="Delete note"
        onCancel={() => setConfirmDelete(false)}
      >
        <p>Are you sure you want to delete this note?</p>
        <div className="modal-actions">
          <button className="btn btn--error" onClick={handleDelete}>
            Yes, delete
          </button>
          <button
            className="btn btn--link"
            onClick={() => setConfirmDelete(false)}
          >
            Cancel
          </button>
        </div>
      </Modal>
    </div>
  );
}
