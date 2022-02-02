
#NoEnv  ; Recommended for performance and compatibility with future AutoHotkey releases.
; #Warn  ; Enable warnings to assist with detecting common errors.
SendMode Input  ; Recommended for new scripts due to its superior speed and reliability.
SetWorkingDir %A_ScriptDir%  ; Ensures a consistent starting directory.

#SingleInstance, Force

; ---------------------------------
; AUTOCOMPLETES
; ---------------------------------
::yout::{bs 1}https://www.youtube.com/feed/subscriptions{enter}
::mail::{bs 1}https://mail.google.com/mail/u/0/{enter}
::twit::{bs 1}https://twitter.com/home{enter}
::nox::Run, nox
; ---------------------------------
; Easy Window Dragging (requires XP/2k/NT)
; ---------------------------------
; Normally, a window can only be dragged by clicking on its title bar.
; This script extends that so that any point inside a window can be dragged.
; To activate this mode, hold down middle mouse button while
; clicking, then drag the window to a new position.

; Note: You can optionally release CapsLock or the middle mouse button after
; pressing down the mouse button rather than holding it down the whole time.
; This script requires v1.0.25+.

~MButton & LButton::
CoordMode, Mouse  ; Switch to screen/absolute coordinates.
MouseGetPos, EWD_MouseStartX, EWD_MouseStartY, EWD_MouseWin
WinGetPos, EWD_OriginalPosX, EWD_OriginalPosY,,, ahk_id %EWD_MouseWin%
WinGet, EWD_WinState, MinMax, ahk_id %EWD_MouseWin% 
if EWD_WinState = 0  ; Only if the window isn't maximized 
	SetTimer, EWD_WatchMouse, 10 ; Track the mouse as the user drags it.
return

EWD_WatchMouse:
GetKeyState, EWD_LButtonState, LButton, P
if EWD_LButtonState = U  ; Button has been released, so drag is complete.
{
	SetTimer, EWD_WatchMouse, Off
	return
}
GetKeyState, EWD_EscapeState, Escape, P
if EWD_EscapeState = D  ; Escape has been pressed, so drag is cancelled.
{
	SetTimer, EWD_WatchMouse, Off
	WinMove, ahk_id %EWD_MouseWin%,, %EWD_OriginalPosX%, %EWD_OriginalPosY%
	return
}
; Otherwise, reposition the window to match the change in mouse coordinates
; caused by the user having dragged the mouse:
CoordMode, Mouse
MouseGetPos, EWD_MouseX, EWD_MouseY
WinGetPos, EWD_WinX, EWD_WinY,,, ahk_id %EWD_MouseWin%
SetWinDelay, -1   ; Makes the below move faster/smoother.
WinMove, ahk_id %EWD_MouseWin%,, EWD_WinX + EWD_MouseX - EWD_MouseStartX, EWD_WinY + EWD_MouseY - EWD_MouseStartY
EWD_MouseStartX := EWD_MouseX  ; Update for the next timer-call to this subroutine.
EWD_MouseStartY := EWD_MouseY
return

; ---------------------------------
; Autoclicker (alt + z)
; ---------------------------------

#maxThreadsPerHotkey, 2
setKeyDelay, 50, 50
setMouseDelay, 50
autoClick := 0

!z::
	autoClick := !autoClick
	while (autoClick = 1)
	{
		click
		sleep 50
	}
return


; ---------------------------------
; Autorun (alt + w)
; ---------------------------------
!w::
	autoClick := !autoClick
	while (autoClick = 1)
	{
		Send {w down}
	}
	Send {w up}
return
; ---------------------------------
; MOUSE BUTTONS - save mouse location / goto mouse location
; ---------------------------------
XButton1::
MouseMove %MouseX%, %MouseY%
return

XButton2::
MouseGetPos MouseX, MouseY
return

; ---------------------------------
; SCRLOCK OFF
; ---------------------------------
#If GetKeyState("ScrollLock","T")
1::
Winset, Alwaysontop, , A
return
#If

; ---------------------------------
; NUMLOCK OFF
; ---------------------------------
; Script Edit / Save
NumpadIns::
SetTitleMatchMode, 2
if not WinExist("startup.ahk")
{
	Edit
}
else
{
	Send ^s
	Reload
	Send ^w
}
return

; Conditional Action
NumpadEnd:: ; 1
SetTitleMatchMode, 2
if WinActive("- Youtube") {
	Send ^+j
	Sleep 2000
	Send {tab}
	Sleep 500
	Send document.getElementsByTagName("video")[0].playbackRate = 3
	Sleep 500
	Send {enter}
	Sleep 500
	Send ^+j
}
else if WinActive("- Brave")
{
	InputBox, tabs, How many image tabs to download?
	if tabs is digit
	{
		loop %tabs%
		{
			Send ^s
			Sleep 1000
			Send {enter}
			Sleep 1000
			Send ^w
			Sleep 250
		}
	}
}
return

NumpadDown:: ; 2
return

NumpadPgDn:: ; 3

return
NumpadLeft:: ; 4
return

NumpadClear:: ; 5
WinGetTitle, Title, A
WinGetClass, Class, A
WinGetPos, WinX, WinY, WinW, WinH, A
MouseGetPos, MouX, MouY
MsgBox, The active window is titled "%Title%" with class "%Class%" and is located at X:"%WinX%" Y:"%WinY%" with size "%WinW%" by "%WinH%". Mouse is at "%MouX%","%MouY%".
return

NumpadRight:: ; 6
return

NumpadHome:: ; 7
return

NumpadUp:: ; 8
return

NumpadPgUp:: ; 9
return

; Move Discord and NOX into place
NumpadDel:: ; decimal
WinGetActiveTitle, Title
WinMove, %Title%,, 1913, 0, 1192, 1057
SetTitleMatchMode, 3
if (WinActive("ahk_class Shell_TrayWnd"))
{
	WinHide, ahk_class Shell_TrayWnd
	Send !{Esc}
}
else
{
	WinShow ahk_class Shell_TrayWnd
	WinActivate ahk_class Shell_TrayWnd
}
SysGet, Mon1, Monitor, 1 ; Right Monitor
SysGet, Mon2, Monitor, 2 ; Left Monitor
Mon1XHalf := Floor(Mon1Right / 2)
Mon1YHalf := Floor(Mon1Bottom / 2)
Mon2XHalf := Floor(Mon2Right / 2)
Mon2YHalf := Floor(Mon2Bottom / 2)
WinMove, NoxPlayer, , %Mon2XHalf%, %Mon2YHalf%, %Mon2XHalf%, %Mon2YHalf%
WinClose, ahk_class Qt5QWindowToolSaveBits
WinMove, ahk_class CabinetWClass, , 0, %Mon2YHalf%, %Mon2XHalf%,%Mon2YHalf%
SetTitleMatchMode, 2
WinMove, Discord, , 0, 0, %Mon2XHalf%, %Mon2YHalf%
WinMove, Picture in picture, , %Mon2XHalf%, 0, %Mon2XHalf%, %Mon2YHalf%
WinMove, gmail.com, , 3091, 0, 516, 1057
WinMove, Twitch, , 0, %Mon2YHalf%, %Mon2XHalf%, %Mon2YHalf%
WinMove, YouTube, , 0, %Mon2YHalf%, %Mon2XHalf%, %Mon2YHalf%
WinMove MINGW64, , 0, %Mon2YHalf%, %Mon2XHalf%, %Mon2YHalf%
return

; Go to %appdata% in explorer and open the discordcanary folder. Then open settings.json and you can add 2 settings: 'MIN_WIDTH' and 'MIN_HEIGHT' to override the default minimums discord imposes on you.

#If !GetKeyState("NumLock","T")
NumpadAdd::
if WinActive("NoxPlayer")
{
	send ^2
}
else if WinActive("VLC")
{
	Send ]]]]]
}
else
{
	Send {Volume_Up 2}
}
return

NumpadSub::
if WinActive("NoxPlayer")
{
	send ^3
}
else if WinActive("VLC")
{
	send [[[[[
}
else
{
	Send {Volume_Down 2}
}
return

; Audio Controls

NumpadDiv::
Run, ms-settings:apps-volume
return

^NumpadDiv::
SoundPlay, C:\Users\Caek\Music\AutoHotKey\Roundabout.mp3
return

; not working
^NumpadMult::
Random, rand, 1, 1968
select := Format("{:04}", rand, -3)
song = "C:\Users\Caek\Music\AutoHotKey\Hibiki\%select%.mp3"
song := RegExReplace(song, "\r\n$","")
;MsgBox, %song%
SoundPlay, %song%, WAIT
return

NumpadMult::
return

NumpadEnter::
Winset, Alwaysontop, , A
return
#If

