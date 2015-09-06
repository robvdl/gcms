import React from 'react';
import AdminContent from './AdminContent';
import AdminToolbar from './AdminToolbar';

/**
 * Admin is the root component for the admin page.
 */
class Admin extends React.Component {
  displayName: 'Admin'

  render() {
    return (
      <div className="admin-container">
        <AdminToolbar />
        <AdminContent />
      </div>
    )
  }
}

export default Admin;
