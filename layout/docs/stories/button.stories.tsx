import type { Meta, StoryObj } from "@storybook/react";
import { Button } from "@noah/ui/button";
import "@noah/ui/styles.css";

const meta: Meta<typeof Button> = {
  component: Button,
  argTypes: {
    type: {
      control: { type: "radio" },
      options: ["button", "submit", "reset"],
    },
  },
};

export default meta;

type Story = StoryObj<typeof Button>;

export const Primary: Story = {
  render: ({ children, ...props }) => <Button {...props}>{children}</Button>,
  name: "Button",
  args: {
    children: "Hello",
    type: "button",
    style: {
      border: "1px solid gray",
      padding: 10,
      borderRadius: 10,
    },
  },
};
