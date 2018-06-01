export interface Streamer {
  id: number;
  tag: string;
  active: boolean;
}

export interface ButtonStyle {
  button_id: string;
  button_color: string;
  button_text: string;
  button_link: string;
  is_hiding: boolean;
}