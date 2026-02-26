import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Home from '@pages/home';
import BlogPosts from '@pages/blogPosts';
import BlogPost from '@pages/blogPost';
import Dashboard from '@pages/dashboard';
import Login from '@pages/login';
import Signup from '@pages/signup';

export default function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/blog" element={<BlogPosts />} />
      <Route path="/blog/:slug" element={<BlogPost />} />
      <Route path="/dashboard" element={<Dashboard />} />
      <Route path="/login" element={<Login />} />
      <Route path="/signup" element={<Signup />} />
    </Routes>
  );
}