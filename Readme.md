# GOLD

```
          _____                   _______                   _____            _____          
         /\    \                 /::\    \                 /\    \          /\    \         
        /::\    \               /::::\    \               /::\____\        /::\    \        
       /::::\    \             /::::::\    \             /:::/    /       /::::\    \       
      /::::::\    \           /::::::::\    \           /:::/    /       /::::::\    \      
     /:::/\:::\    \         /:::/~~\:::\    \         /:::/    /       /:::/\:::\    \     
    /:::/  \:::\    \       /:::/    \:::\    \       /:::/    /       /:::/  \:::\    \    
   /:::/    \:::\    \     /:::/    / \:::\    \     /:::/    /       /:::/    \:::\    \   
  /:::/    / \:::\    \   /:::/____/   \:::\____\   /:::/    /       /:::/    / \:::\    \  
 /:::/    /   \:::\ ___\ |:::|    |     |:::|    | /:::/    /       /:::/    /   \:::\ ___\ 
/:::/____/  ___\:::|    ||:::|____|     |:::|    |/:::/____/       /:::/____/     \:::|    |
\:::\    \ /\  /:::|____| \:::\    \   /:::/    / \:::\    \       \:::\    \     /:::|____|
 \:::\    /::\ \::/    /   \:::\    \ /:::/    /   \:::\    \       \:::\    \   /:::/    / 
  \:::\   \:::\ \/____/     \:::\    /:::/    /     \:::\    \       \:::\    \ /:::/    /  
   \:::\   \:::\____\        \:::\__/:::/    /       \:::\    \       \:::\    /:::/    /   
    \:::\  /:::/    /         \::::::::/    /         \:::\    \       \:::\  /:::/    /    
     \:::\/:::/    /           \::::::/    /           \:::\    \       \:::\/:::/    /     
      \::::::/    /             \::::/    /             \:::\    \       \::::::/    /      
       \::::/    /               \::/____/               \:::\____\       \::::/    /       
        \::/____/                 ~~                      \::/    /        \::/____/        
                                                           \/____/          ~~              
                                                                                            
```
![Logo](Example/demo/logo.png)

## GOLD

**GOLD** is a Go graphics library, built on top of [Ebiten](https://ebiten.org/), for creating GUIs and 2D games.

### Main Features

- Graphical widgets (buttons, sliders, etc.)
- Sprite and image management
- Mouse/keyboard event system
- Easy to integrate into your Go projects

### Example Usage

```go
import . "github.com/vad3l/gold/library/graphics/gui"

btn := NewButton(Point{100, 40}, Point{10, 10}, "Click me")
btn.Draw(screen)
btn.Input()
if btn.Execute {
    // Do something!
}
```

### Installation

> ⚠️ The repository uses a custom import path. To use it locally:
```sh
git clone https://github.com/vad3l/gold.git
# then import locally in your projects
```

### Dependencies

- [Ebiten](https://github.com/hajimehoshi/ebiten)
- [golang/freetype](https://github.com/golang/freetype)

---

**Made by vad3l**