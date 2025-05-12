import React, { useState } from "react";
import api from "../shared/api";
import { Link, useNavigate } from "react-router-dom";
import "./SignUpPage.css"

export default function SignupPage() {
  const [form, setForm] = useState({ name: "", username: "", password: "" });
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  const handleChange = (e) => setForm({ ...form, [e.target.name]: e.target.value });

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      await api.post("/auth/sign-up", form);
      navigate("/login");
    } catch (err) {
      setError(err?.response?.data?.message || err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <section className="signup-page">
      <div className="signup-card">
        <h2 className="signup-title">Create an account</h2>

        <form className="signup-form" onSubmit={handleSubmit}>
          <div className="signup-field">
            <label htmlFor="name" className="signup-label">
              Full name
            </label>
            <input
              id="name"
              name="name"
              type="text"
              className="signup-input"
              placeholder="Jane Doe"
              value={form.name}
              onChange={handleChange}
              required
            />
          </div>

          <div className="signup-field">
            <label htmlFor="username" className="signup-label">
              Username
            </label>
            <input
              id="username"
              name="username"
              type="text"
              className="signup-input"
              placeholder="jane_doe"
              value={form.username}
              onChange={handleChange}
              required
            />
          </div>

          <div className="signup-field">
            <label htmlFor="password" className="signup-label">
              Password
            </label>
            <input
              id="password"
              name="password"
              type="password"
              className="signup-input"
              placeholder="••••••••"
              value={form.password}
              onChange={handleChange}
              required
            />
          </div>

          <button type="submit" className="signup-button" disabled={loading}>
            {loading ? "Signing up…" : "Sign Up"}
          </button>

          {error && <div className="signup-error">{error}</div>}
        </form>

        <p className="signup-footer">
          Already have an account? <Link to="/login">Sign in</Link>
        </p>
      </div>
    </section>
  );
}