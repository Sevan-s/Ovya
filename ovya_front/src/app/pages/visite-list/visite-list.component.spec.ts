import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VisiteListComponent } from './visite-list.component';

describe('VisiteListComponent', () => {
  let component: VisiteListComponent;
  let fixture: ComponentFixture<VisiteListComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [VisiteListComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(VisiteListComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
