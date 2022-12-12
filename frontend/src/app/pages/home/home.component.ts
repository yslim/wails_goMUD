import {AfterViewInit, Component, ElementRef, OnDestroy, OnInit, ViewChild} from '@angular/core';
import {Terminal} from "xterm";
import {FitAddon} from "xterm-addon-fit";
import {debounceTime, fromEvent, map, startWith, Subscription} from "rxjs";
import {Send} from "../../../wailsjs/go/service/MudService";
import {EventsOn, EventsEmit} from "../../../wailsjs/runtime";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements AfterViewInit, OnInit, OnDestroy {
  @ViewChild('term') terminalElement!: ElementRef;
  inputBuffer = "";
  terminal: Terminal | undefined;
  termFit: FitAddon | undefined;
  windowResizeSubs: Subscription | undefined;

  ngOnInit(): void {
    this.windowResizeSubs = windowSizeObserver(500).subscribe(({width, height}) => {
      this.termFit?.fit();
    })

    EventsOn("OnMessage", (message: string) => {
      this.terminal?.write(message);
    });
  }

  ngOnDestroy(): void {
    this.windowResizeSubs?.unsubscribe();
  }

  ngAfterViewInit(): void {
    // init terminal & addons
    this.terminal = new Terminal({
      allowProposedApi: true,
      fontFamily: "Gulimche",
      fontSize: 17,
    });
    this.termFit = new FitAddon();
    this.terminal.loadAddon(this.termFit);

    // open terminal in DOM
    this.terminal.open(this.terminalElement.nativeElement);
    this.termFit.fit();

    this.terminal.onData((data) => {
      this.terminal?.write(data);
    });

    EventsEmit("OnReady");
  }

  onKeyDown($event: KeyboardEvent) {
    if ($event.key === "Enter") {
      Send(this.inputBuffer);
      this.inputBuffer = "";
    }
  }
}

function windowSizeObserver(dTime = 300) {
  return fromEvent(window, 'resize').pipe(
    debounceTime(dTime),
    map(event => {
      const window = event.target as Window;
      return {width: window.innerWidth, height: window.innerHeight}
    }),
    startWith({width: window.innerWidth, height: window.innerHeight})
  );
}
