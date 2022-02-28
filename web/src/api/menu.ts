import request from '@/utils/request';

export const getMenus = () => {
  return request({
    url: '/system/menus',
    method: 'get',
  });
};
