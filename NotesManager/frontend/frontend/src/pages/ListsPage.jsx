import React, { useEffect, useState } from "react";
import api from "../shared/api";
import { useNavigate } from "react-router-dom";
import ListCard from "../components/ListCard";
import Modal from "../components/Modal";
import "./ListsPage.css";

export default function ListsPage() {
  const [lists, setLists] = useState([]);
  const [newTitle, setNewTitle] = useState("");
  const [newDesc, setNewDesc] = useState("");
  const [query, setQuery] = useState("");
  const [sortOrder, setSortOrder] = useState("desc");
  const [viewMode, setViewMode] = useState("grid");
  const [editList, setEditList] = useState(null);
  const [deleteList, setDeleteList] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetchLists();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [sortOrder]);

  const fetchLists = async () => {
    setLoading(true);
    try {
      const res = await api.get("api/lists", {
        params: { q: query.trim(), sort: sortOrder },
      });
      setLists(res.data.data ?? []);
    } catch (err) {
      setError(err?.response?.data?.message || err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleAdd = async () => {
    if (!newTitle.trim()) return;
    await api.post("api/lists", {
      title: newTitle.trim(),
      description: newDesc.trim(),
    });
    setNewTitle("");
    setNewDesc("");
    fetchLists();
  };

  const openEditModal = (list) => setEditList({ ...list });
  const handleSaveEdit = async () => {
    if (!editList) return;
    await api.put(`api/lists/${editList.id}`, {
      title: editList.title,
      description: editList.description,
    });
    setEditList(null);
    fetchLists();
  };

  useEffect(() => { console.log('editList >>', editList); }, [editList]);

  const openDeleteModal = (list) => setDeleteList(list);
  const handleDeleteConfirm = async () => {
    if (!deleteList) return;
    await api.delete(`api/lists/${deleteList.id}`);
    setDeleteList(null);
    fetchLists();
  };

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  const toggleSort = () => setSortOrder((p) => (p === "asc" ? "desc" : "asc"));
  const toggleView = () => setViewMode((v) => (v === "grid" ? "list" : "grid"));

  return (
    <section className="lists-page">
      <header className="lists-header">
        <h1 className="lists-title">Notes Manager</h1>
        <button className="btn btn--link" onClick={handleLogout}>
          Logout
        </button>
      </header>

      <main className="lists-main">
        <div className="toolbar">
          <div className="toolbar-section toolbar-section--search">
            <input
              type="text"
              className="input input--search"
              placeholder="Search lists…"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
            />
            <button className="btn btn--primary" onClick={fetchLists}>
              Search
            </button>
          </div>

          <div className="toolbar-section toolbar-section--actions">
            <button className="btn btn--secondary" onClick={toggleSort}>
              Sort: {sortOrder === "asc" ? "Oldest" : "Newest"}
            </button>
            <button className="btn btn--secondary" onClick={toggleView}>
              View: {viewMode === "grid" ? "List" : "Grid"}
            </button>
          </div>
        </div>

        <div className="add-form">
          <input
            type="text"
            className="input"
            placeholder="New list title"
            value={newTitle}
            onChange={(e) => setNewTitle(e.target.value)}
          />
          <input
            type="text"
            className="input"
            placeholder="Description (optional)"
            value={newDesc}
            onChange={(e) => setNewDesc(e.target.value)}
          />
          <button className="btn btn--primary" onClick={handleAdd}>
            Add List
          </button>
        </div>

        {loading ? (
          <p className="state state--loading">Loading…</p>
        ) : error ? (
          <p className="state state--error">{error}</p>
        ) : lists.length === 0 ? (
          <p className="state state--muted">No lists found.</p>
        ) : viewMode === "grid" ? (
          <div className="lists-grid">
            {lists.map((list) => (
              <ListCard
                key={list.id}
                list={list}
                onEdit={() => openEditModal(list)}
                onDelete={() => openDeleteModal(list)}
              />
            ))}
          </div>
        ) : (
          <div className="lists-list">
            {lists.map((list) => (
              <ListCard
                key={list.id}
                list={list}
                onEdit={() => openEditModal(list)}
                onDelete={() => openDeleteModal(list)}
              />
            ))}
          </div>
        )}
      </main>

      <Modal isOpen={!!editList} title="Edit List" onCancel={() => setEditList(null)}>
        <div className="modal-form">
          <input
            className="input-title"
            value={editList?.title || ""}
            onChange={(e) => setEditList((f) => ({ ...f, title: e.target.value }))}
            placeholder="Title"
          />
          <input
            className="input-desc"
            value={editList?.description || ""}
            onChange={(e) =>
              setEditList((f) => ({ ...f, description: e.target.value }))
            }
            placeholder="Description"
          />
          <button className="btn btn--primary" id="form-save-btn" onClick={handleSaveEdit}>
            Save
          </button>
        </div>
      </Modal>

      <Modal isOpen={!!deleteList} title="Confirm Delete" onCancel={() => setDeleteList(null)}>
        <p>Delete list "{deleteList?.title}"?</p>
        <div className="modal-actions">
          <button className="btn btn--error" onClick={handleDeleteConfirm}>
            Yes, Delete
          </button>
          <button className="btn btn--link" onClick={() => setDeleteList(null)}>
            Cancel
          </button>
        </div>
      </Modal>
    </section>
  );
}