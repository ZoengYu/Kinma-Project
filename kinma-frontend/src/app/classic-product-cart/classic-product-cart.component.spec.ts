import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ClassicProductCartComponent } from './classic-product-cart.component';

describe('ClassicProductCartComponent', () => {
  let component: ClassicProductCartComponent;
  let fixture: ComponentFixture<ClassicProductCartComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ClassicProductCartComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ClassicProductCartComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
