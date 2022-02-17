import { Logout, User, UserProfile } from '@vicons/carbon';
import { NButton, NDropdown, NIcon } from 'naive-ui';
import { defineComponent, h } from 'vue';

const renderIcon = (icon: any) => {
  return () => {
    return h(NIcon, null, {
      default: () => h(icon),
    });
  };
};

const options = [
  {
    label: '用户资料',
    key: 'profile',
    icon: renderIcon(UserProfile),
  },
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(Logout),
  },
];

export default defineComponent({
  setup() {
    return () => (
      <NDropdown options={options}>
        <NButton
          text
          v-slots={{
            icon: () => <User />,
            default: () => <span>个人中心</span>,
          }}
        />
      </NDropdown>
    );
  },
});
