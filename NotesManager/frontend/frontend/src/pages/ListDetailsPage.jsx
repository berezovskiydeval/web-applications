import React, { useEffect, useState, useRef } from "react";
import api from "../shared/api";
import { useParams, useNavigate } from "react-router-dom";
import Modal from "../components/Modal";
import "./ListDetailsPage.css";

export default function ListDetailsPage() {
  const { listId } = useParams();
  const navigate = useNavigate();

  const [query, setQuery] = useState("");
  const [sortOrder, setSortOrder] = useState("desc");

  const [list, setList] = useState(null);
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");
  const [pinned, setPinned] = useState(false);

  const [openMenuId, setOpenMenuId] = useState(null);
  const [confirmDelete, setConfirmDelete] = useState(null);
  const menuRef = useRef(null);

  useEffect(() => {
    if (listId) loadAll();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [listId, sortOrder]);

  useEffect(() => {
    const handleClickOutside = (e) => {
      if (menuRef.current && !menuRef.current.contains(e.target)) {
        setOpenMenuId(null);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  async function loadAll() {
    setLoading(true);
    setError(null);
    try {
      const { data: listWrap } = await api.get(`/api/lists/${listId}`);
      setList(listWrap.data ?? listWrap);

      const { data: itemsWrap } = await api.get(`/api/lists/${listId}/items`, {
        params: { q: query.trim(), sort: sortOrder },
      });
      setItems(itemsWrap.data ?? []);
    } catch (e) {
      setError(e.message);
    } finally {
      setLoading(false);
    }
  }

  async function handleAddNote() {
    if (!title.trim() || !content.trim()) return;
    try {
      await api.post(`api/lists/${listId}/items`, {
        title: title.trim(),
        content: content.trim(),
        pinned,
      });
      setTitle("");
      setContent("");
      setPinned(false);
      loadAll();
    } catch (e) {
      setError(e.message);
    }
  }

  async function handleTogglePin(item) {
    try {
      await api.put(`api/items/${item.id}`, {
        title: item.title,
        content: item.content,
        pinned: !item.pinned,
      });
      loadAll();
    } catch (e) {
      setError(e.message);
    }
  }

  async function handleDeleteConfirmed() {
    if (!confirmDelete) return;
    try {
      await api.delete(`api/items/${confirmDelete.id}`);
      setConfirmDelete(null);
      loadAll();
    } catch (e) {
      setError(e.message);
    }
  }

  const toggleSort = () => setSortOrder((prev) => (prev === "asc" ? "desc" : "asc"));

  if (loading) return <p>Loading…</p>;
  if (error) return <p className="text text--error">{error}</p>;
  if (!list) return <p className="text text--muted">List not found.</p>;

  return (
    <section className="details-page">
      <header className="details-header">
        <button className="btn btn--link" onClick={() => navigate("/lists")}>← Back</button>
        <h2 className="details-title">{list.title}</h2>
      </header>

      <main className="details-main">
        {list.description && <p className="text text--muted">{list.description}</p>}

        <div className="toolbar">
          <input
            type="text"
            className="input input--search"
            placeholder="Search notes…"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
          />
          <button className="btn btn--primary" onClick={loadAll}>Search</button>
          <button className="btn btn--secondary" onClick={toggleSort}>
            Sort: {sortOrder === "asc" ? "Oldest" : "Newest"}
          </button>
        </div>

        <div className="note-form">
          <input
            className="input-title-note"
            placeholder="Title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <textarea
            className="input textarea"
            placeholder="Content"
            rows={4}
            value={content}
            onChange={(e) => setContent(e.target.value)}
          />
          <label className="form-checkbox">
            <input type="checkbox" checked={pinned} onChange={(e) => setPinned(e.target.checked)} /> Pinned
          </label>
          <button className="btn btn--primary" onClick={handleAddNote}>Add Note</button>
        </div>

        {items.length === 0 ? (
          <p className="text text--muted">No notes found.</p>
        ) : (
          <ul className="notes-grid">
            {items.map((item) => (
              <li key={item.id} className="note-card">
                <div className="note-meta">
                  <span>{new Date(item.created_at).toLocaleString()}</span>
                  <button
                    className="note-menu-btn"
                    onClick={() => setOpenMenuId((id) => (id === item.id ? null : item.id))}
                    aria-label="Actions"
                  >
                    ⋮
                  </button>
                  {openMenuId === item.id && (
                    <div className="note-menu" ref={menuRef}>
                      <button onClick={() => { handleTogglePin(item); setOpenMenuId(null); }}>
                        {item.pinned ? "Unpin" : "Pin"}
                      </button>
                      <button onClick={() => { navigate(`/lists/${listId}/items/${item.id}/edit`); }}>
                        Edit
                      </button>
                      <button onClick={() => { setConfirmDelete(item); setOpenMenuId(null); }}>
                        Delete
                      </button>
                    </div>
                  )}
                </div>

                <h3 className="note-title" onClick={() => navigate(`/lists/${listId}/items/${item.id}`)}>
                  {item.title}
                </h3>
                <p className="note-content" onClick={() => navigate(`/lists/${listId}/items/${item.id}`)}>
                  {item.content}
                </p>
                {item.pinned && <span className="badge badge--info">Pinned</span>}
              </li>
            ))}
          </ul>
        )}
      </main>

      <Modal
        isOpen={!!confirmDelete}
        title="Delete note"
        onCancel={() => setConfirmDelete(null)}
      >
        <p>Delete note "{confirmDelete?.title}"?</p>
        <div className="modal-actions">
          <button className="btn btn--error" onClick={handleDeleteConfirmed}>Yes, delete</button>
          <button className="btn btn--link" onClick={() => setConfirmDelete(null)}>Cancel</button>
        </div>
      </Modal>
    </section>
  )}