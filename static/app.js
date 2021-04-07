/*
 * @Date: 2021-04-07 10:04:32
 * @LastEditors: viletyy
 * @LastEditTime: 2021-04-07 10:04:33
 * @FilePath: /egg/static/app.js
 */
document.addEventListener('DOMContentLoaded', () => {
  (document.querySelectorAll('.notification .delete') || []).forEach(($delete) => {
    const $notification = $delete.parentNode;

    $delete.addEventListener('click', () => {
      $notification.parentNode.removeChild($notification);
    });
  });
});