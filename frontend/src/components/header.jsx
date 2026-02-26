import React from 'react';
import { Link } from 'react-router-dom';
import * as Icons from 'lucide-react';

const Icon = ({ name, ...props }) => {
  const LucideIcon = Icons?.[name] || Icons.HelpCircle;
  return <LucideIcon {...props} />;
};

export default function Header() {
  return (
    <header className="sticky top-0 z-50 bg-white/80 backdrop-blur-xl border-b border-blue-100/30 shadow-soft">
      <div className="container mx-auto px-4 md:px-6">
        <div className="flex items-center justify-between h-20">
          <Link to="/" className="flex items-center gap-3 group">
            <div className="w-10 h-10 rounded-lg bg-gradient-to-br from-blue-500 to-purple-500 flex items-center justify-center shadow-soft group-hover:shadow-lg transition-all duration-300">
              <Icon name="PenTool" className="w-5 h-5 text-white" />
            </div>
            <span className="text-xl font-bold text-slate-800 tracking-tight">BlogSphere</span>
          </Link>
          <nav className="hidden md:flex items-center gap-8">
            <Link to="/login" className="text-slate-600 hover:text-blue-600 transition-colors duration-300">Login</Link>
            <Link to="/signup" className="text-slate-600 hover:text-blue-600 transition-colors duration-300">Sign Up</Link>
            <Link to="/dashboard" className="text-slate-600 hover:text-blue-600 transition-colors duration-300">Dashboard</Link>
          </nav>
          <button className="md:hidden text-slate-600 hover:text-blue-600 transition-colors duration-300">
            <Icon name="Menu" className="w-6 h-6" />
          </button>
        </div>
      </div>
    </header>
  );
}