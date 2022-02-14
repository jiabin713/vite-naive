import { h } from 'vue';
import { BookOutline } from '@vicons/ionicons5';
import { NIcon, NLayoutSider, NMenu } from 'naive-ui';
import { MenuMixedOption } from 'naive-ui/lib/menu/src/interface';
import { defineComponent } from 'vue';

function renderIcon(icon: any) {
  return () => h(NIcon, null, { default: () => h(icon) });
}

const menuOptions: MenuMixedOption[] = [
  {
    label: '且听风吟',
    key: 'hear-the-wind-sing',
    icon: renderIcon(BookOutline),
  },
  {
    label: '1973年的弹珠玩具',
    key: 'pinball-1973',
    icon: renderIcon(BookOutline),
    disabled: true,
    children: [
      {
        label: '鼠',
        key: 'rat',
      },
    ],
  },
  {
    label: '寻羊冒险记',
    key: 'a-wild-sheep-chase',
    disabled: true,
    icon: renderIcon(BookOutline),
  },
  {
    label: '舞，舞，舞',
    key: 'dance-dance-dance',
    icon: renderIcon(BookOutline),
    children: [
      {
        type: 'group',
        label: '人物',
        key: 'people',
        children: [
          {
            label: '叙事者',
            key: 'narrator',
            icon: renderIcon(BookOutline),
          },
          {
            label: '羊男',
            key: 'sheep-man',
            icon: renderIcon(BookOutline),
          },
        ],
      },
      {
        label: '饮品',
        key: 'beverage',
        icon: renderIcon(BookOutline),
        children: [
          {
            label: '威士忌',
            key: 'whisky',
          },
        ],
      },
      {
        label: '食物',
        key: 'food',
        children: [
          {
            label: '三明治',
            key: 'sandwich',
          },
        ],
      },
      {
        label: '过去增多，未来减少',
        key: 'the-past-increases-the-future-recedes',
      },
    ],
  },
  {
    label: '舞，舞，舞',
    key: 'dance-dance-dance',
    icon: renderIcon(BookOutline),
    children: [
      {
        type: 'group',
        label: '人物',
        key: 'people',
        children: [
          {
            label: '叙事者',
            key: 'narrator',
            icon: renderIcon(BookOutline),
          },
          {
            label: '羊男',
            key: 'sheep-man',
            icon: renderIcon(BookOutline),
          },
        ],
      },
      {
        label: '饮品',
        key: 'beverage',
        icon: renderIcon(BookOutline),
        children: [
          {
            label: '威士忌',
            key: 'whisky',
          },
        ],
      },
      {
        label: '食物',
        key: 'food',
        children: [
          {
            label: '三明治',
            key: 'sandwich',
          },
        ],
      },
      {
        label: '过去增多，未来减少',
        key: 'the-past-increases-the-future-recedes',
      },
    ],
  },
  {
    label: '舞，舞，舞',
    key: 'dance-dance-dance',
    icon: renderIcon(BookOutline),
    children: [
      {
        type: 'group',
        label: '人物',
        key: 'people',
        children: [
          {
            label: '叙事者',
            key: 'narrator',
            icon: renderIcon(BookOutline),
          },
          {
            label: '羊男',
            key: 'sheep-man',
            icon: renderIcon(BookOutline),
          },
        ],
      },
      {
        label: '饮品',
        key: 'beverage',
        icon: renderIcon(BookOutline),
        children: [
          {
            label: '威士忌',
            key: 'whisky',
          },
        ],
      },
      {
        label: '食物',
        key: 'food',
        children: [
          {
            label: '三明治',
            key: 'sandwich',
          },
        ],
      },
      {
        label: '过去增多，未来减少',
        key: 'the-past-increases-the-future-recedes',
      },
    ],
  },
  {
    label: '舞，舞，舞',
    key: 'dance-dance-dance',
    icon: renderIcon(BookOutline),
    children: [
      {
        type: 'group',
        label: '人物',
        key: 'people',
        children: [
          {
            label: '叙事者',
            key: 'narrator',
            icon: renderIcon(BookOutline),
          },
          {
            label: '羊男',
            key: 'sheep-man',
            icon: renderIcon(BookOutline),
          },
        ],
      },
      {
        label: '饮品',
        key: 'beverage',
        icon: renderIcon(BookOutline),
        children: [
          {
            label: '威士忌',
            key: 'whisky',
          },
        ],
      },
      {
        label: '食物',
        key: 'food',
        children: [
          {
            label: '三明治',
            key: 'sandwich',
          },
        ],
      },
      {
        label: '过去增多，未来减少',
        key: 'the-past-increases-the-future-recedes',
      },
    ],
  },
];

export default defineComponent({
  setup() {
    return () => (
      <NLayoutSider bordered showTrigger nativeScrollbar={false}>
        <NMenu
          // collapsedWidth={48}
          // collapsedIconSize={40}
          options={menuOptions}
        />
      </NLayoutSider>
    );
  },
});
