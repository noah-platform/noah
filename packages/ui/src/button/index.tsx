export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
}

export function Button({ children, ...other }: ButtonProps): JSX.Element {
  return (
    <button className="ui-bg-slate-600 ui-text-white" type="button" {...other}>
      {children}
    </button>
  );
}

Button.displayName = "Button";
